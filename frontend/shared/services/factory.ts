import { AxiosInstance } from 'axios'

type ServiceCreator<T> = (client: AxiosInstance) => T

export class ServiceFactory {
	constructor(private client: AxiosInstance) {}

	create<T>(creator: ServiceCreator<T>): T {
		return creator(this.client)
	}
}
