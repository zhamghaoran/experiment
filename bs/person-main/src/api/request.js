import axios from "axios";

let request = axios.create({
	baseURL: "/api",
	timeout: 3 * 1000,
});
request.interceptors.request.use(
	config => {
		config.headers.Accept = "application/json";
		return config;
	},
	error => {
		return Promise.reject(error);
	}
);

request.interceptors.response.use(
	response => {
		return response;
	},
	error => {
		return Promise.reject(error);
	}
);

export default request;
