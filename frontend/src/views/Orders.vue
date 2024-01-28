<template>
  <div>
    <el-container>
      <el-header>
        <h2>Orders</h2>
      </el-header>
      <el-main>
        <el-table :data="orders" style="width: 100%" stripe>
          <el-table-column label="Order Name" prop="order_name"></el-table-column>
          <el-table-column label="Company Name" prop="company_name"></el-table-column>
          <el-table-column label="Customer Name" prop="customer_name"></el-table-column>
          <el-table-column label="Order Date" prop="order_date"></el-table-column>
          <el-table-column label="Total Amount" prop="total_amount"></el-table-column>
          <el-table-column label="Delivered Quantity" prop="delivered_quantity"></el-table-column>
        </el-table>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { ElTable, ElTableColumn } from 'element-plus';
import axios from 'axios';

export default {
  components: {
    ElTable,
    ElTableColumn,
  },
  setup() {
    const orders = ref([]);

    onMounted(async () => {
      try {
        // Replace 'YOUR_API_ENDPOINT' with the actual API endpoint
        const response = await axios.get('http://localhost:5000/orders');
        orders.value = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    });

    return {
      orders,
    };
  },
};
</script>

<style>
/* Add your custom styles here */
</style>
