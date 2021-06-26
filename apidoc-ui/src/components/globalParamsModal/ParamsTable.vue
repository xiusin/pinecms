<template>
  <div style="height:342px">
    <Table
      :columns="columns"
      size="small"
      rowKey="id"
      :bordered="true"
      :pagination="false"
      :data-source="currentData"
      :scroll="tableScroll"
    >
      <template slot="editRowKey" slot-scope="text, record">
        <TableInput
          :data="text"
          placeholder="请输入参数名"
          @change="onCellChange(record, 'key', $event)"
        />
      </template>
      <template slot="editRowValue" slot-scope="text, record">
        <TableInput
          :data="text"
          placeholder="请输入参数值"
          @change="onCellChange(record, 'value', $event)"
        />
      </template>
      <template slot="editRowDesc" slot-scope="text, record">
        <TableInput
          :data="text"
          placeholder="请输入参数说明"
          @change="onCellChange(record, 'desc', $event)"
        />
      </template>
      <template slot="actionRow" slot-scope="text, record">
        <a-button
          icon="delete"
          type="danger"
          ghost
          @click="deleteRow(record)"
        ></a-button>
      </template>
    </Table>
    <a-button style="margin-top:10px" @click="addRow">+ 添加参数</a-button>
  </div>
</template>

<script>
import { Table, Button } from "ant-design-vue";
import TableInput from "@/utils/Input";
export default {
  components: {
    Table,
    TableInput,
    [Button.name]: Button
  },
  props: {
    device: {
      type: String,
      default: "xl"
    },
    data: {
      type: Array,
      default: () => []
    }
  },

  data() {
    return {
      columns: [
        {
          title: "Key",
          dataIndex: "key",
          width: 150,
          scopedSlots: { customRender: "editRowKey" }
        },
        {
          title: "Value",
          dataIndex: "value",
          width: this.device == "mobile" ? 150 : 350,
          scopedSlots: { customRender: "editRowValue" }
        },
        {
          title: "说明",
          dataIndex: "desc",
          scopedSlots: { customRender: "editRowDesc" }
        },
        {
          title: "操作",
          dataIndex: "id",
          width: 50,
          scopedSlots: { customRender: "actionRow" }
        }
      ],
      currentData: [],
      tableScroll: {
        x: "700px",
        y: "260px"
      }
    };
  },
  created() {
    this.currentData = this.data;
  },
  methods: {
    onCellChange(row, field, value) {
      row[field] = value;
    },
    addRow() {
      this.currentData.push({
        id: new Date().getTime() + Math.ceil(Math.random() * 100),
        key: "",
        value: "",
        desc: ""
      });
    },
    deleteRow(row) {
      this.currentData = this.currentData.filter(p => p.id !== row.id);
    },
    getData() {
      return this.currentData;
    }
  }
};
</script>
