import request from "./request";
export const testAPI = {
	getAll() {
		return request({
			method: "get",
			url: "/show",
		});
	},
	addPerosn(data) {
		return request({
			method: "post",
			url: "/add",
			data: {
				姓名: data.姓名,
				电话: data.电话,
				地址: data.地址,
			},
		});
	},
	delPerson() {
		return request({
			method: "post",
			url: "/delete",
		});
	},
	modify(data) {
		console.log({...data});
		return request({
			method: "post",
			url: "/modify",
			data: {...data},
		});
	},
};
