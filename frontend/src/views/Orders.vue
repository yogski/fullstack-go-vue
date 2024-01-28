<template>
  <div>
    <el-container>
      <el-main>
        <h2>Orders</h2>
        <el-input style="width: 40%;"
          v-model="searchQuery"
          placeholder="Search..."
          @keyup.enter="searchOrders"
        ></el-input>
        <el-button type="primary" @click="searchOrders">Search</el-button>
      </el-main>
    </el-container>
  </div>
  <div>
    <el-container>
      <el-main>
        <el-date-picker
          v-model="dateRange"
          type="daterange"
          range-separator="to"
          start-placeholder="Start Date"
          end-placeholder="End Date"
          @change="searchOrders"
        ></el-date-picker>
      </el-main>
    </el-container>
  </div>
  <div>
    <el-container>
      <el-main>
        <el-table :data="pagedOrders" style="width: 100%" stripe>
          <el-table-column label="Order Name" prop="order_name"></el-table-column>
          <el-table-column label="Company Name" prop="company_name"></el-table-column>
          <el-table-column label="Customer Name" prop="customer_name"></el-table-column>
          <el-table-column label="Order Date" prop="order_date"></el-table-column>
          <el-table-column label="Total Amount" prop="total_amount"></el-table-column>
          <el-table-column label="Delivered Quantity" prop="delivered_quantity"></el-table-column>
        </el-table>
        <p>pagination should be between this .../</p>
      </el-main>
    </el-container>
  </div>
  <div>
    <el-container>
      <el-main>
        <el-pagination
          @current-change="handlePageChange"
          :current-page="currentPage"
          :page-sizes="[5, 10, 20]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="orders.length"
        ></el-pagination>
        <p>and this... :/</p>
      </el-main>
    </el-container>
  </div>

</template>

<script>
import { ref, onMounted, computed } from 'vue';
import { ElTable, ElTableColumn, ElPagination, ElInput, ElButton, ElDatePicker } from 'element-plus';
import axios from 'axios';

export default {
  components: {
    ElTable,
    ElTableColumn,
    ElPagination,
    ElInput,
    ElButton,
    ElDatePicker,
  },
  setup() {
    const orders = ref([]);
    const currentPage = ref(1);
    const pageSize = ref(5);
    const searchQuery = ref('');
    const dateRange = ref([]);

    // Calculate the start index based on the current page and page size
    const startIndex = computed(() => (currentPage.value - 1) * pageSize.value);

    // Calculate the end index based on the start index and page size
    const endIndex = computed(() => startIndex.value + pageSize.value);

    // Get a slice of the orders array based on the start and end indices
    const pagedOrders = computed(() => orders.value.slice(startIndex.value, endIndex.value));

    onMounted(async () => {
      await fetchOrders();
    });

    const fetchOrders = async () => {
      try {
        // Replace 'YOUR_API_ENDPOINT' with the actual API endpoint
        const API_ENDPOINT = 'http://localhost:5000'
        const response = await axios.get(`${API_ENDPOINT}/orders`, { 
          params: {
            q: searchQuery.value,
            start_date: dateRange.value[0],
            end_date: dateRange.value[1],
          },
        });
        orders.value = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    const handlePageChange = (newPage) => {
      currentPage.value = newPage;
    };

    const searchOrders = async () => {
      currentPage.value = 1; // Reset page to 1 when performing a new search
      await fetchOrders();
    };

    return {
      orders,
      currentPage,
      pageSize,
      searchQuery,
      dateRange,
      pagedOrders,
      handlePageChange,
      searchOrders,
    };
  },
};
</script>

<style>
/* Add your custom styles here */
</style>
