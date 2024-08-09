
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="交易类型" prop="transactionType"><el-select v-model="queryParams.transactionType"
                                               placeholder="交易记录交易类型" clearable size="small">
                                        <el-option
                                                v-for="dict in transactionTypeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="交易日期" prop="transactionDate"><el-input v-model="queryParams.transactionDate" placeholder="请输入交易日期" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="备注" prop="remarks"><el-input v-model="queryParams.remarks" placeholder="请输入备注" clearable
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
                                v-permisaction="['admin:sysStockTransactionRecords:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:sysStockTransactionRecords:edit']"
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
                                v-permisaction="['admin:sysStockTransactionRecords:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="sysStockTransactionRecordsList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="商品编码" align="center" prop="productId" :formatter="productIdFormat" width="100">
                                <template slot-scope="scope">
                                    {{ productIdFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="交易类型" align="center" prop="transactionType"
                                                 :formatter="transactionTypeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ transactionTypeFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="数量" align="center" prop="quantity"
                                                 :show-overflow-tooltip="true"/><el-table-column label="交易日期" align="center" prop="transactionDate"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.transactionDate) }}</span>
                                    </template>
                                </el-table-column><el-table-column label="备注" align="center" prop="remarks"
                                                 :show-overflow-tooltip="true"/>
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
                             v-permisaction="['admin:sysStockTransactionRecords:edit']"
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
                              v-permisaction="['admin:sysStockTransactionRecords:remove']"
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
                                        <el-select v-model="form.productId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in productIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="交易类型" prop="transactionType">
                                        <el-select v-model="form.transactionType"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in transactionTypeOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="数量" prop="quantity">
                                        <el-input v-model="form.quantity" placeholder="数量"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="交易日期" prop="transactionDate">
                                        <el-date-picker
                                                    v-model="form.transactionDate"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                                    <el-form-item label="备注" prop="remarks">
                                        <el-input v-model="form.remarks" placeholder="备注"
                                                      />
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
    import {addSysStockTransactionRecords, delSysStockTransactionRecords, getSysStockTransactionRecords, listSysStockTransactionRecords, updateSysStockTransactionRecords} from '@/api/admin/sys-stock-transaction-records'
    
    import {listSysStockProducts } from '@/api/admin/sys-stock-products'
    export default {
        name: 'SysStockTransactionRecords',
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
                sysStockTransactionRecordsList: [],
                transactionTypeOptions: [],
                // 关系表类型
                productIdOptions :[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    transactionType:undefined,
                    transactionDate:undefined,
                    remarks:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {transactionType:  [ {required: true, message: '交易类型不能为空', trigger: 'blur'} ],
                transactionDate:  [ {required: true, message: '交易日期不能为空', trigger: 'blur'} ],
                remarks:  [ {required: true, message: '备注不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getSysStockProductsItems()
            this.getDicts('sys_stock_status').then(response => {
                this.transactionTypeOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listSysStockTransactionRecords(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.sysStockTransactionRecordsList = response.data.list
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
                transactionType: undefined,
                quantity: undefined,
                transactionDate: undefined,
                remarks: undefined,
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
            transactionTypeFormat(row) {
                return this.selectDictLabel(this.transactionTypeOptions, row.transactionType)
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
                this.title = '添加交易记录'
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
                getSysStockTransactionRecords(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改交易记录'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateSysStockTransactionRecords(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addSysStockTransactionRecords(this.form).then(response => {
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
                      return delSysStockTransactionRecords( { 'ids': Ids })
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
