package repository

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type operatorMap = map[string]string
type fieldsMap = map[string]string

// supported operators, those preceded by __
var inlineOperators = []string{"in", "between"}

type Parser struct {
	fields fieldsMap
	model  interface{}
	Filter *Filter
	log    *slog.Logger
}

func NewParser(model interface{}, log *slog.Logger) *Parser {
	return &Parser{
		model: model,
		Filter: &Filter{
			Offset: 0,
			Order:  "asc",
			Sort:   "id",
		},
		log: log,
	}
}

// use the model to get the json field names and their corresponding db field names
func (p *Parser) getFieldsFor(model interface{}) fieldsMap {
	fields := make(map[string]string)

	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		var jsonFieldName string
		jsonField := v.Type().Field(i).Tag.Get("json")

		// this field may have additional info like omitempty so we check for the comma
		if strings.Contains(jsonField, ",") {
			jsonFieldName = strings.Split(jsonField, ",")[0]
		} else {
			jsonFieldName = jsonField
		}
		fields[jsonFieldName] = v.Type().Field(i).Tag.Get("db")

		// if the db tag is empty, we use the json field name
		if fields[jsonFieldName] == "" {
			fields[jsonFieldName] = jsonFieldName
		}
	}
	return fields
}

// parses the limit field. Must have a value greater than 0
func (p *Parser) parseLimit(params url.Values) error {
	if params.Get("limit") != "" {
		limit, err := strconv.Atoi(params.Get("limit"))
		if err != nil {
			p.log.Error(err.Error())
			return errors.New("limit must be an integer")
		}
		if limit < 0 {
			return errors.New("limit must be greater than 0")
		}
		p.Filter.Limit = limit
	}
	return nil
}

// parses the offset field. Must have a value greater than or equal to 0
func (p *Parser) parseOffset(params url.Values) error {
	if params.Get("offset") != "" {
		offset, err := strconv.Atoi(params.Get("offset"))
		if err != nil {
			p.log.Error(err.Error())
			return errors.New("offset must be an integer")
		}
		if offset < 0 {
			return errors.New("offset must be greater than or equal to 0")
		}
		p.Filter.Offset = offset
	}
	return nil
}

// handles sort=field or sort=field.asc or sort=field.desc
func (p *Parser) parseSort(params url.Values) error {
	if params.Get("sort") != "" {
		sort := params.Get("sort")
		if strings.Contains(sort, ".") {
			splits := strings.Split(params.Get("sort"), ".")
			if len(splits) == 2 && splits[1] == "" {
				return errors.New("missing order")
			}
			field, order := splits[0], splits[1]
			dbField, ok := p.fields[field]
			if !ok {
				return errors.New("invalid sort field")
			}
			if order != "asc" && order != "desc" {
				return errors.New("invalid sort order")
			}
			p.Filter.Sort = dbField
			p.Filter.Order = order
		} else {
			dbField, ok := p.fields[sort]
			if !ok {
				return errors.New("invalid sort field")
			}
			p.Filter.Sort = dbField
		}
	}
	return nil
}

// limited operator support
// allow for <, >, <=, >=, ~=, __ and we add a default of = if no operator is provided
func (p *Parser) parseOperator(k string, v []string) (string, string, string, error) {
	var operator string
	var dbField string
	var val string
	var ok bool

	if strings.HasSuffix(k, "<") || strings.HasSuffix(k, ">") || strings.HasSuffix(k, "~") {
		field := k[:len(k)-1]
		dbField, ok = p.fields[field]
		if !ok {
			return "", "", "", errors.New("invalid filter field")
		}
		operator = k[len(k)-1:] + "="
		val = v[0]
	} else {
		if strings.Contains(k, "<") {
			splits := strings.Split(k, "<")
			field := splits[0]
			dbField, ok = p.fields[field]
			if !ok {
				return "", "", "", errors.New("invalid filter field")
			}
			operator = "<"
			val = splits[1]
		}
		if strings.Contains(k, ">") {
			splits := strings.Split(k, ">")
			field := splits[0]
			dbField, ok = p.fields[field]
			if !ok {
				return "", "", "", errors.New("invalid filter field")
			}
			operator = ">"
			val = splits[1]
		}
		if strings.Contains(k, "__") { // this supports multiple operators preceded by __
			splits := strings.Split(k, "__")
			field := splits[0]
			dbField, ok = p.fields[field]
			if !ok {
				return "", "", "", errors.New("invalid filter field")
			}
			operator = splits[1]
			if !slices.Contains(inlineOperators, operator) {
				return "", "", "", errors.New("invalid operator: " + operator)
			}
			val = v[0]
		}
	}
	return dbField, operator, val, nil
}

// parses the filter fields
func (p *Parser) parseFilter(params url.Values) error {
	filterMap := make(map[string]interface{})
	operatorMap := make(operatorMap)
	var dbField string
	var ok bool
	for k, v := range params {
		if k != "limit" && k != "offset" && k != "sort" {
			if strings.Contains(k, "<") || strings.Contains(k, ">") || strings.Contains(k, "~") || strings.Contains(k, "__") {
				dbField, operator, val, err := p.parseOperator(k, v)
				if err != nil {
					return err
				}
				operatorMap[dbField] = operator

				if slices.Contains(inlineOperators, operator) {
					filterMap[dbField], err = getMultipleValues(operator, val)
					if err != nil {
						return err
					}
				} else {
					filterMap[dbField] = val
				}
			} else {
				dbField, ok = p.fields[k]
				operatorMap[dbField] = "="
				filterMap[dbField] = v[0]
				if !ok {
					return fmt.Errorf("invalid filter field: %v", k)
				}
			}
		}
	}
	p.Filter.Fields = filterMap
	p.Filter.Operators = operatorMap
	return nil
}

// returns list of values for inline operators, separated by pipe
func getMultipleValues(operator string, val string) ([]string, error) {
	var values []string

	// creating slice of strings and trimming whitespaces
	inValues := strings.Split(val, "|")
	for vIndex, vVal := range inValues {
		inValues[vIndex] = strings.TrimSpace(vVal)
	}
	values = inValues

	if operator == "between" && len(values) != 2 {
		return nil, errors.New("values quantity must be two")
	}

	return values, nil
}

// parses the url params and returns a Filter struct
// we allow for the following params:
// limit, offset, sort, sort.<direction>, field<operator>value
func (p *Parser) Parse(params url.Values) (*Filter, error) {
	p.fields = p.getFieldsFor(p.model)
	err := p.parseLimit(params)
	if err != nil {
		return nil, err
	}
	err = p.parseOffset(params)
	if err != nil {
		return nil, err
	}
	err = p.parseSort(params)
	if err != nil {
		return nil, err
	}
	err = p.parseFilter(params)
	if err != nil {
		return nil, err
	}

	return p.Filter, nil
}