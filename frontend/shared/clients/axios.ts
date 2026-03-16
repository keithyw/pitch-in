import axios, { AxiosInstance, AxiosError } from 'axios'
import { RefreshResponse } from '@pitch-in/shared/types'

export interface AxiosOptions {
	baseUrl?: string
	getToken: () => string | null
	getRefreshToken: () => string | null
	onRefreshSuccess: (res: RefreshResponse) => void
	onLogout: () => void
	refreshUrl: string
	authUrls: string[]
	redirectUrl?: string
}

export const createAxiosClient = (opts: AxiosOptions): AxiosInstance => {
	const client = axios.create({
		baseURL: opts.baseUrl,
		headers: {
			'Content-Type': 'application/json',
			Accept: 'application/json',
		},
		timeout: 10000,
	})
	let isRefreshing = false
	let failedQueue: any[] = []

	const processQueue = (
		error: AxiosError | null,
		token: string | null = null,
	) => {
		failedQueue.forEach((p) => (error ? p.reject(error) : p.resolve(token)))
		failedQueue = []
	}

	client.interceptors.request.use((config) => {
		const token = opts.getToken()
		if (token) config.headers.Authorization = `Bearer ${token}`
		return config
	})

	client.interceptors.response.use(
		(res) => res,
		async (error: AxiosError) => {
			const req: any = error.config
			if (error.response?.status !== 401 || req._retry)
				return Promise.reject(error)
			// if the url is token/login url
			if (opts.authUrls.some((url) => req.url?.includes(url)))
				return Promise.reject(error)
			if (isRefreshing) {
				return new Promise((resolve, reject) => {
					failedQueue.push({ resolve, reject })
				})
			}

			req._retry = true
			isRefreshing = true

			try {
				const refresh = opts.getRefreshToken()
				if (!refresh) {
					throw new Error('No refresh token found')
				}
				const postRes = await axios.post(`${opts.baseUrl}${opts.refreshUrl}`, {
					refresh,
				})
				opts.onRefreshSuccess(postRes.data)
				processQueue(null, postRes.data.access) // not sure if this will be the structure in terms of the refresh url payload

				// need to set heades and new refresh token
				req.headers = {
					...req.headers,
					Authorization: `Bearer ${refresh}`,
				}
				return client(req)
			} catch (err: unknown) {
				processQueue(err as AxiosError)
				opts.onLogout()
				// need to set windows login href, print to console and return reject promise
				if (opts.redirectUrl) window.location.href = opts.redirectUrl
				console.error('Refresh token failed,redirecting to login: ', err)
				return Promise.reject(err)
			} finally {
				isRefreshing = false
			}
		},
	)

	return client
}
