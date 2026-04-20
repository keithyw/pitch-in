export class AIServiceException extends Error {
	status: number

	constructor(message: string, status: number) {
		super(message)
		this.status = status
		Object.setPrototypeOf(this, AIServiceException.prototype)
	}
}

export interface PromptHintComponentProps {
	onHandleSubmit: () => void
}

export interface IdentifierResponse {
	name: string
	description: string
	tags: string[]
}

export interface IngestionRequest {
	name: string
	description: string
	tags: string[]
}

export interface IngestionResponse {}
