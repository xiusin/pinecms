<template>
    <cl-dialog title="模板配置" :close-on-click-modal="false" v-model="visible">
        <el-form :model="dataForm" :rules="dataRule" ref="dataForm" label-width="100px" size="mini">
            <el-form-item label="标题" prop="title" size="mini">
                <el-input v-model="dataForm.title" placeholder="标题" size="mini"></el-input>
            </el-form-item>
            <el-form-item label="链接" prop="url" size="mini">
                <el-input v-model="dataForm.url" placeholder="跳转链接" size="mini"></el-input>
            </el-form-item>
            <div>
                <el-form-item label="小程序appid" prop="miniprogram.appid" size="mini">
                    <el-input v-model="dataForm.miniprogram.appid" placeholder="小程序appid" size="mini"></el-input>
                </el-form-item>
                <el-form-item label="小程序路径" prop="miniprogram.pagePath">
                    <el-input v-model="dataForm.miniprogram.pagePath" placeholder="小程序pagePath" size="mini"></el-input>
                </el-form-item>
            </div>
            <el-row>
                <el-col :span="16">
                    <el-form-item label="模版名称" prop="name">
                        <el-input v-model="dataForm.name" placeholder="模版名称" size="mini"></el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="8">
                    <el-form-item label="有效" prop="status">
                        <el-switch v-model="dataForm.status" placeholder="是否有效" :active-value="true" :inactive-value="false" size="mini"></el-switch>
                    </el-form-item>
                </el-col>
            </el-row>
            <div class="form-group-area">

				<el-card class="box-card" :closable="false" shadow="hover">
					<div slot="header" class="clearfix">
						<el-alert show-icon type="warning" :closable="false" effect="dark">消息填充数据，请对照模板内容填写</el-alert>
					</div>
					<div style="padding: 5px; font-size: 12px; background-color: #dfe1e5"><code><pre>{{dataForm.content}}</pre></code></div>
				</el-card>
				<div style="height: 15px;"></div>
                <el-row v-for="(item,index) in dataForm.data" :key="item.name">
                    <el-col :span="16">
                        <el-form-item :label="item.name" :prop="'data.'+index+'.value'" :rules="[{required: true,message: '填充内容不得为空', trigger: 'blur' }]">
                            <el-input type="textarea" autosize rows="1" v-model="item.value" size="mini" placeholder="填充内容"  ></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="8">
                        <el-form-item label="颜色" >
                            <el-input type="color" v-model="item.color" placeholder="颜色" size="mini"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </div>
			<el-form-item  size="mini">
			</el-form-item>
		</el-form>

        <span slot="footer" class="dialog-footer" style="padding-top: 10px ;">
            <el-button @click="visible = false" size="mini">取消</el-button>
            <el-button type="primary" @click="dataFormSubmit()" size="mini">确定</el-button>
        </span>
    </cl-dialog>
</template>

<script>
export default {
    data() {
        return {
            visible: false,
            dataForm: {
                id: 0,
                templateId: '',
                title: '',
                data: [],
                url: '',
                miniprogram:{appid:'',pagePath:''},
                content: '',
                status: true,
                name: ''
            },
            dataRule: {
                title: [
                    { required: true, message: '标题不能为空', trigger: 'blur' }
                ],
                data: [
                    { required: true, message: '内容不能为空', trigger: 'blur' }
                ],
                name: [
                    { required: true, message: '模版名称不能为空', trigger: 'blur' }
                ]
            }
        }
    },
    methods: {
        init(id) {
        	this.visible = true;
            this.dataForm.id = id || 0
            this.visible = true
            this.$nextTick(() => {
                this.$refs['dataForm'].resetFields()
                if (this.dataForm.id) {
                	this.service.wechat.template.info({"id": this.dataForm.id}).then((data) => {
						this.transformTemplate(data)
					})
                }
            })
        },
        transformTemplate(template){
        	console.log(template)
            if(!template.miniprogram){
            	template.miniprogram={appid:'',pagePath:''}
			}
            if(template.data instanceof Array) {//已经配置过了，直接读取
                this.dataForm =  template
                return
            }

            template.data=[]
            let keysArray = template.content.match(/\{\{(\w*)\.DATA\}\}/g) || [] //["{{first.DATA}}"]
            keysArray.map(item=>{
                name=item.replace('{{','').replace('.DATA}}','')
                template.data.push({"name":name,"value":"",color:"#000000"})
            })
            this.dataForm = template
        },
        // 表单提交
        dataFormSubmit() {
            this.$refs['dataForm'].validate((valid, msg) => {
                if (valid) {
                	if (!this.dataForm.id) {
						this.service.wechat.template.add(this.dataForm).then((data) => {
							this.$message.success(data);
						})
					} else {
						this.service.wechat.template.update(this.dataForm).then((data) => {
							this.$message.success(data);
						}).catch((e) => {
							this.$message.error(e);
						});
					}
                }
            })
        }
    }
}
</script>
<style scoped>
.form-group-area{
    border:1px dotted gray;
	padding: 10px;
}
</style>
