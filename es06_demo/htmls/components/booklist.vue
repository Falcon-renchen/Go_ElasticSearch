<template>
    <div>
      <div>
          <dl class="sdt">
              <dt>
                  出版社
              </dt>
              <dd><a href="">清华大学出版社</a></dd>
          </dl>
      </div>
        <div>
            <el-table
                    :data="bookList"
                    style="width: 100%">
                <el-table-column
                        prop="BookId"
                        label="ID"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="BookName"
                        label="商品名称"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="price1"
                        label="价格">
                </el-table-column>
                <el-table-column
                        prop="KindName"
                        label="分类">
                </el-table-column>
            </el-table>
        </div>

    </div>

</template>
<script>
   module.exports ={
       data(){
           return {
               bookList: []
           }
       }
   }
</script>
<style>
  .sdt{margin: 0 auto ;width:90%;border-radius: 5px;}
  .sdt dt{width:100%;display: block;color:#3A7B43;font-size:16px;font-weight: bold}
  .sdt dd{float:left;margin:0 auto;text-indent:1em}
  .sdt dd a{color: #3a8ee6}
  .sdt dd a:hover{background: #eb5975;color:#fff}
</style>