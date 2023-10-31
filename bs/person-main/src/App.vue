<script setup>
	import { ref, onMounted } from "vue";
	import { testAPI } from "./api/test";
	import { ElMessage } from "element-plus";

	let dataList = ref("");
	let dialogVisible = ref(false);
	//state : 1添加联系人 0：修改
	let state = ref(1);
	let personFormRef = ref();
	let personForm = ref({
		id: "",
		姓名: "",
		电话: "",
		地址: "",
	});
	let personFormRules = ref({
		// id:[{}]
		姓名: [{ required: true, message: "请输入名字", trigger: "change" }],
		电话: [
			{ required: true, message: "请输入电话号", trigger: "change" },
			{
				validator: (rule, value, callback) => {
					if (
						/^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/.test(
							value
						) === false
					) {
						callback(new Error("手机号码格式不正确"));
					} else {
						callback();
					}
				},
				trigger: "blur",
			},
		],
		地址: [{ required: true, message: "请输入地址", trigger: "change" }],
	});
	//获取所有联系人
	async function getAllPerson() {
		let res = await testAPI.getAll();
		console.log(res);
		dataList.value = res.data.message;
	}
	onMounted(() => {
		getAllPerson();
	});
	//修改按钮的回调
	function modifyHandler(item) {
		state.value = 0;
		personForm.value = { ...item };
		dialogVisible.value = true;
	}
	//添加按钮的回调
	function addHandler() {
		state.value = 1;
		if (personFormRef.value) {
			personFormRef.value.resetFields();
		}
		dialogVisible.value = true;
		personForm.value = {};
	}
	//修改添加的确认回调
	async function confirmHandler() {
		if (state.value == 1) {
			if (!personFormRef.value) {
				return;
			}
			personFormRef.value.validate(async valid => {
				if (valid) {
					let res = await testAPI.addPerosn(personForm.value);
					if (res.code === 200) {
						ElMessage({
							type: "success",
							message: "添加成功",
						});
					}
					getAllPerson();
					dialogVisible.value = false;
				} else {
					return false;
				}
			});
		} else if (state.value == 0) {
			if (!personFormRef.value) {
				return;
			}
			personFormRef.value.validate(async valid => {
				if (valid) {
					let res = await testAPI.modify(personForm.value);
					if (res.code === 200) {
						ElMessage({
							type: "success",
							message: "修改成功",
						});
					}
					getAllPerson();
					dialogVisible.value = false;
				} else {
					return false;
				}
			});
		}
	}
	//删除操作
	async function delHandler(item) {
		let res = await testAPI.delPerson(item.id);
		if (res.code === 200) {
			ElMessage({
				type: "success",
				message: "删除成功",
			});
		}
	}
</script>

<template>
	<div class="test">
		<div>
			<el-card class="card">
				<template #header>
					<div class="card-header">
						<span>联系人管理</span>
						<el-button type="primary" @click="addHandler()">添加</el-button>
					</div>
				</template>
				<div class="card-body">
					<el-table :data="dataList" style="width: 100%" :highlight-current-row="true">
						<el-table-column prop="姓名" label="姓名" width="80" align="center" />
						<el-table-column prop="电话" label="电话" width="150" align="center" />
						<el-table-column prop="地址" label="地址" width="150" align="center" />
						<el-table-column label="操作" width="200" align="center">
							<template #default="{ row }">
								<div>
									<el-button type="primary" link @click="modifyHandler(row)">
										修改
									</el-button>
									<el-popconfirm
										title="Are you sure to delete this?"
										@confirm="delHandler(row)"
									>
										<template #reference>
											<el-button type="danger" link> 删除 </el-button>
										</template>
									</el-popconfirm>
								</div>
							</template>
						</el-table-column>
					</el-table>
				</div>
			</el-card>
		</div>
		<el-dialog :title="state === 1 ? '添加联系人' : '修改联系人'" v-model="dialogVisible">
			<div>
				<el-form
					ref="personFormRef"
					:model="personForm"
					:rules="personFormRules"
					label-width="80px"
				>
					<el-form-item label="姓名" prop="姓名">
						<el-input v-model="personForm.姓名"></el-input>
					</el-form-item>
					<el-form-item label="电话" prop="电话">
						<el-input v-model="personForm.电话"></el-input>
					</el-form-item>
					<el-form-item label="地址" prop="地址">
						<el-input v-model="personForm.地址"></el-input>
					</el-form-item>
				</el-form>
			</div>
			<div slot="footer">
				<el-button @click="dialogVisible = false">取 消</el-button>
				<el-button type="primary" @click="confirmHandler">确 定</el-button>
			</div>
		</el-dialog>
	</div>
</template>

<style scoped lang="scss">
	.test {
		display: flex;
		flex-direction: column;
		align-items: center;
		padding: 100px 0;
		.card {
			width: 700px;
			.card-header {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 0 30px;
			}
		}
	}
</style>
