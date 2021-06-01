<template>
    <div>
        <div>
            <dl class="sdt">
                <dt>
                    出版社
                </dt>
                <dd v-for="item in pressList"><a href="">{{item}}</a></dd>
            </dl>
            <dl class="sdt">
                <dt>
                    全文检索
                </dt>
                <dd class="row">
                    <el-input  v-model="searchModel.book_name"   class="search" placeholder="请输入内容"></el-input>
                    <el-button type="primary"  @click="search" >搜索</el-button>
                </dd>
            </dl>
        </div>
        <div>
            <el-table
                    :data="bookList"
                    style="width: 100%">
                <el-table-column
                        prop="BookID"
                        label="ID"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="BookName"
                        label="商品名称"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="BookIntr"
                        label="商品简介"
                        width="280">
                </el-table-column>
                <el-table-column
                        prop="BookPrice1"
                        label="价格">
                </el-table-column>
                <el-table-column
                        prop="BookPress"
                        label="出版社">
                </el-table-column>
                <el-table-column
                        prop="BookAuthor"
                        label="作者">
                </el-table-column>
            </el-table>
        </div>

    </div>

</template>
<script>
    module.exports ={
        data(){
            return {
                bookList: [],
                pressList: [],
                searchModel: {
                    book_name: ''
                }
            }
        },
        created(){
            this.loadPress()  // 加载出版社列表
        },
        methods: {
            async loadPress(){
                const response = await axios.get("/helper/press")
                const { result } = response.data
                this.pressList = result
            },
            async search(){
                const response = await axios.post("/books/search",this.searchModel)
                const { result } = response.data
                this.bookList = result
            }
        }
    }
</script>
<style>
    .sdt{margin: 30px auto;width:90%;border-radius: 5px;display: block;float:left;margin-left: 50px}
    .sdt dt{width:100%;display: block;color:#3A7B43;font-size:16px;font-weight: bold}
    .sdt dd{float:left;margin:0 auto;text-indent:1em}
    .sdt .row{width:100%;}
    .sdt dd .search{width:50%;}
    .sdt dd a{color: #3a8ee6}
    .sdt dd a:hover{background: #eb5975;color:#fff}
</style>