
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="商品编码" prop="productId"><el-select v-model="queryParams.productId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in productIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="数量" prop="quantity"><el-input v-model="queryParams.quantity" placeholder="请输入数量" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="最后更新时间" prop="lastUpdated"><el-input v-model="queryParams.lastUpdated" placeholder="请输入最后更新时间" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:sysStockInventory:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:sysStockInventory:edit']"
                                type="success"
                                icon="el-icon-edit"
                                size="mini"
                                :disabled="single"
                                @click="handleUpdate"
                        >修改
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:sysStockInventory:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="sysStockInventoryList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="商品编码" align="center" prop="productId" :formatter="productIdFormat" width="100">
                                <template slot-scope="scope">
                                    {{ productIdFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="数量" align="center" prop="quantity"
                                                 :show-overflow-tooltip="true"/><el-table-column label="最后更新时间" align="center" prop="lastUpdated"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.lastUpdated) }}</span>
                                    </template>
                                </el-table-column>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                         <el-popconfirm
                           class="delete-popconfirm"
                           title="确认要修改吗?"
                           confirm-button-text="修改"
                           @confirm="handleUpdate(scope.row)"
                         >
                           <el-button
                             slot="reference"
                             v-permisaction="['admin:sysStockInventory:edit']"
                             size="mini"
                             type="text"
                             icon="el-icon-edit"
                           >修改
                           </el-button>
                         </el-popconfirm>
                         <el-popconfirm
                            class="delete-popconfirm"
                            title="确认要删除吗?"
                            confirm-button-text="删除"
                            @confirm="handleDelete(scope.row)"
                         >
                            <el-button
                              slot="reference"
                              v-permisaction="['admin:sysStockInventory:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>
                        </template>
                    </el-table-column>
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <el-dialog :title="title" :visible.sync="open" width="500px">
                    <el-form ref="form" :model="form" :rules="rules" label-width="80px">
                        
                                    <el-form-item label="商品编码" prop="productId">
                                        <el-input v-model="form.productId" placeholder="商品编码"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="数量" prop="quantity">
                                        <el-input v-model="form.quantity" placeholder="数量"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="最后更新时间" prop="lastUpdated">
                                        <el-date-picker
                                                    v-model="form.lastUpdated"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addSysStockInventory, delSysStockInventory, getSysStockInventory, listSysStockInventory, updateSysStockInventory} from '@/api/admin/sys-stock-inventory'
    
    import {listSysStockProducts } from '@/api/admin/sys-stock-products'
    export default {
        name: 'SysStockInventory',
        components: {
        },
        data() {
            return {
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                sysStockInventoryList: [],
                
                // 关系表类型
                productIdOptions :[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    productId:undefined,
                    quantity:undefined,
                    lastUpdated:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {productId:  [ {required: true, message: '商品编码不能为空', trigger: 'blur'} ],
                quantity:  [ {required: true, message: '数量不能为空', trigger: 'blur'} ],
                lastUpdated:  [ {required: true, message: '最后更新时间不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getSysStockProductsItems()
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listSysStockInventory(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.sysStockInventoryList = response.data.list
                        this.total = response.data.count
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.form = {
                
                id: undefined,
                productId: undefined,
                quantity: undefined,
                lastUpdated: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            productIdFormat(row) {
                return this.selectItemsLabel(this.productIdOptions, row.productId)
            },
            // 关系
            getSysStockProductsItems() {
               this.getItems(listSysStockProducts, undefined).then(res => {
                   this.productIdOptions = this.setItems(res, 'id', 'name')
               })
            },
            // 文件
            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1
                this.getList()
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            /** 新增按钮操作 */
            handleAdd() {
                this.reset()
                this.open = true
                this.title = '添加库存'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const id =
                row.id || this.ids
                getSysStockInventory(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改库存'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateSysStockInventory(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addSysStockInventory(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                var Ids = (row.id && [row.id]) || this.ids

                this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(function () {
                      return delSysStockInventory( { 'ids': Ids })
                }).then((response) => {
                   if (response.code === 200) {
                     this.msgSuccess(response.msg)
                     this.open = false
                     this.getList()
                   } else {
                     this.msgError(response.msg)
                   }
                }).catch(function () {
                })
            }
        }
    }
</script>
