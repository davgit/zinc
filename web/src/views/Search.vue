<template>
  <q-page class="q-pa-md">
    <div class="column">
      <search-bar
        ref="searchBar"
        :data="queryData"
        @updated="queryUpdated"
        @refresh="searchData"
      />
      <div class="row">
        <index-list ref="indexList" :data="indexData" @updated="indexUpdated" />
        <search-list ref="searchList" @updated:fields="updateIndexFields" />
      </div>
    </div>
  </q-page>
</template>

<script>
import { defineComponent, ref } from "vue";

import SearchBar from "../components/search/SearchBar.vue";
import IndexList from "../components/search/IndexList.vue";
import SearchList from "../components/search/SearchList.vue";

export default defineComponent({
  name: "PageSearch",
  components: {
    SearchBar,
    IndexList,
    SearchList,
  },
  setup() {
    const indexData = ref({
      name: "",
      columns: [],
    });
    const queryData = ref({
      query: "",
      time: {
        tab: "relative",
        startDate: "",
        startTime: "",
        endDate: "",
        endTime: "",
        selectedRelativePeriod: "Minutes",
        selectedRelativeValue: 30,
        selectedFullTime: false,
      },
    });

    const searchBar = ref(null);
    const indexList = ref(null);
    const searchList = ref(null);
    const searchData = () => {
      searchList.value.searchData(indexData.value, queryData.value);
    };
    const resetColumns = () => {
      searchList.value.resetColumns(indexData.value);
    };

    const indexUpdated = ({ name, columns }) => {
      if (indexData.value.name != name) {
        indexData.value.name = name;
        indexData.value.columns = columns;
        queryData.value.query = "";
        searchBar.value.setSearchQuery("");
        searchData();
      } else {
        indexData.value.columns = columns;
        resetColumns();
      }
    };

    const queryUpdated = ({ query, time }) => {
      queryData.value.query = query;
      queryData.value.time = time;
      searchData();
    };

    const updateIndexFields = (fields) => {
      indexList.value.setIndexFields(fields);
    };

    return {
      indexData,
      queryData,
      searchBar,
      indexList,
      searchList,
      searchData,
      indexUpdated,
      queryUpdated,
      updateIndexFields,
    };
  },
});
</script>
