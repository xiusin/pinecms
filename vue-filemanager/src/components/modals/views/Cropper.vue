<template >
  <div class="cropper">
    <el-dialog :visible="showCropper" width="85%" :before-close="handleClose">
      <span slot="title" class="title">
        <strong>裁剪</strong>
        <small class="text-truncate">{{ selectedItem.basename }}</small>
      </span>
      <el-row type="flex" style="max-height: 460px">
        <el-col :span="18">
          <img :src="imgSrc" ref="Cropper" :alt="selectedItem.basename" />
        </el-col>
        <el-col :span="6">
          <div class="cropper-preview"></div>
          <div class="cropper-data">
            <el-form label-position="left" :model="ruleForm" :rules="rules" @submit.native.prevent>
              <el-form-item label="X" label-width="50px" prop="x">
                <el-input v-model.number="ruleForm.x"></el-input>
                <span class="unit">px</span>
              </el-form-item>
              <el-form-item label="Y" label-width="50px" prop="y">
                <el-input v-model.number="ruleForm.y"></el-input>
                <span class="unit">px</span>
              </el-form-item>
              <el-form-item label="Width" label-width="50px" prop="width">
                <el-input v-model.number="ruleForm.width"></el-input>
                <span class="unit">px</span>
              </el-form-item>
              <el-form-item label="Height" label-width="50px" prop="height">
                <el-input v-model.number="ruleForm.height"></el-input>
                <span class="unit">px</span>
              </el-form-item>
              <el-form-item label="Rotate" label-width="50px" prop="rotate">
                <el-input v-model.number="ruleForm.rotate"></el-input>
                <span class="unit">deg</span>
              </el-form-item>
              <el-form-item label="ScaleX" label-width="50px" prop="scaleX">
                <el-input v-model.number="ruleForm.scaleX"></el-input>
              </el-form-item>
              <el-form-item label="ScaleY" label-width="50px" prop="scaleY">
                <el-input v-model.number="ruleForm.x"></el-input>
              </el-form-item>
            </el-form>
            <el-button @click="setData()" title="应用" type="info" class="check">
              <i class="fas fa-check" />
            </el-button>
          </div>
        </el-col>
      </el-row>

      <span slot="footer" class="dialog-footer">
        <el-row type="flex" justify="space-between">
          <div style="display:flex;">
            <div class="btn-group" role="group" aria-label="Scale">
              <el-button @click="cropMove(-10, 0)" type="info">
                <i class="fas fa-arrow-left" />
              </el-button>
              <el-button @click="cropMove(10, 0)" type="info">
                <i class="fas fa-arrow-right" />
              </el-button>
              <el-button @click="cropMove(0, -10)" type="info">
                <i class="fas fa-arrow-up" />
              </el-button>
              <el-button @click="cropMove(0, 10)" type="info">
                <i class="fas fa-arrow-down" />
              </el-button>
            </div>
            <div class="btn-group" role="group" aria-label="Scale">
              <el-button @click="cropScaleX()" type="info">
                <i class="fas fa-arrows-alt-h" />
              </el-button>
              <el-button @click="cropScaleY()" type="info">
                <i class="fas fa-arrows-alt-v" />
              </el-button>
            </div>
            <div class="btn-group" role="group" aria-label="Rotate">
              <el-button @click="cropRotate(-45)" type="info">
                <i class="fas fa-undo" />
              </el-button>
              <el-button @click="cropRotate(45)" type="info">
                <i class="fas fa-redo" />
              </el-button>
            </div>
            <div class="btn-group" role="group" aria-label="Rotate">
              <el-button @click="cropZoom(0.1)" type="info">
                <i class="fas fa-search-plus" />
              </el-button>
              <el-button @click="cropZoom(-0.1)" type="info">
                <i class="fas fa-search-minus" />
              </el-button>
            </div>
            <el-button @click="cropReset()" title="重置" class="info">
              <i class="fas fa-sync-alt" />
            </el-button>
            <el-button @click="cropSave()" title="保存" type="danger" class="save">
              <i class="far fa-save" />
            </el-button>
          </div>
          <div class="back">
            <el-button size="small" @click="showModal('Preview')" class="light">返回</el-button>
          </div>
        </el-row>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import Cropper from "cropper";
import EventBus from "@/eventBus";
export default {
  props: {
    showCropper: { required: true, type: Boolean },
    imgSrc: { required: true, type: String }
    //     maxHeight: { type: Number, required: true }
  },
  data() {
    const ruleArr = val => {
      return [
        { required: true, message: "内容不能为空", trigger: "blur" },
        { type: "number", message: `${val}必须为数字值`, trigger: "blur" },
        { validator: validate, trigger: "blur" }
      ];
    };
    const validate = (rule, value, callback) => {
      if (value) {
        this.setData();
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        height: 0,
        width: 0,
        x: 0,
        y: 0,
        rotate: 0,
        scaleX: 1,
        scaleY: 1
      },
      rules: {
        x: ruleArr("X"),
        y: ruleArr("Y"),
        height: ruleArr("Height"),
        width: ruleArr("Width"),
        rotate: ruleArr("Rotate"),
        scaleX: ruleArr("ScaleX"),
        scaleY: ruleArr("ScaleY")
      },
      cropper: {}
    };
  },
  computed: {
    /**
     * ACL on/off
     */
    acl() {
      return this.$store.state.fm.settings.acl;
    },
    /**
     * 选择文件
     * @returns {*}
     */
    selectedItem() {
      return this.$store.getters["fm/selectedItems"][0];
    },
    /**
     * 计算图片的最大高度
     * @returns {number}
     */
    maxHeight() {
      if (this.$store.state.fm.modal.modalBlockHeight) {
        return this.$store.state.fm.modal.modalBlockHeight - 170;
      }

      return 300;
    },
    itemAuthor() {
      return this.selectedItem.author === this.$store.state.fm.username;
    }
  },

  mounted() {
    this.$nextTick(function() {
      // 设置裁剪器实例
      this.cropper = new Cropper(this.$refs.Cropper, {
        preview: ".cropper-preview",
        crop: e => {
          this.ruleForm.x = Math.round(e.detail.x);
          this.ruleForm.y = Math.round(e.detail.y);
          this.ruleForm.height = Math.round(e.detail.height);
          this.ruleForm.width = Math.round(e.detail.width);
          this.ruleForm.rotate =
            typeof e.detail.rotate !== "undefined" ? e.detail.rotate : "";
          this.ruleForm.scaleX =
            typeof e.detail.scaleX !== "undefined" ? e.detail.scaleX : "";
          this.ruleForm.scaleY =
            typeof e.detail.scaleY !== "undefined" ? e.detail.scaleY : "";
        }
      });
    });
  },
  beforeDestroy() {
    this.cropper.destroy();
  },

  methods: {
    /**
     * 平移
     * @param x轴
     * @param y轴
     */
    cropMove(x = 0, y = 0) {
      this.cropper.move(x, y);
    },
    /**
     * Y轴镜像
     */
    cropScaleY() {
      this.cropper.scale(1, this.cropper.getData().scaleY === 1 ? -1 : 1);
    },

    /**
     * X轴镜像
     */
    cropScaleX() {
      this.cropper.scale(this.cropper.getData().scaleX === 1 ? -1 : 1, 1);
    },
    /**
     * 旋转
     * @param grade
     */
    cropRotate(grade) {
      this.cropper.rotate(grade);
    },

    /**
     * 缩放
     * @param ratio
     */
    cropZoom(ratio) {
      this.cropper.zoom(ratio);
    },
    /**
     * 重置还原
     */
    cropReset() {
      this.cropper.reset();
    },

    /**
     * 从表单中设置数据
     */
    setData() {
      this.cropper.setData({
        x: this.ruleForm.x,
        y: this.ruleForm.y,
        width: this.ruleForm.width,
        height: this.ruleForm.height,
        rotate: this.ruleForm.rotate,
        scaleX: this.ruleForm.scaleX,
        scaleY: this.ruleForm.scaleY
      });
    },
    /**
     * 保存裁剪的图片
     */
    cropSave() {
      if (this.acl && this.selectedItem.acl === 1 && !this.itemAuthor) {
        this.$notify.warning({
          title: "警告",
          message: "没有修改此图片的权限!"
        });
        return;
      }
      this.cropper.getCroppedCanvas().toBlob(
        blob => {
          const formData = new FormData();
          // 添加当前磁盘名称
          formData.append("disk", this.$store.getters["fm/selectedDisk"]);
          // 添加路径
          formData.append("path", this.selectedItem.dirname);
          // 数据、图片名称
          formData.append("file", blob, this.selectedItem.basename);
          this.$store.dispatch("fm/updateFile", formData).then(resp => {
            if (resp.data.result.status === "success") {
              // 关闭裁剪器
              // this.$emit("closeCropper");
              EventBus.$emit("showCropper", false, "");
              this.$store.commit("fm/modal/setModalState", {
                modalName: "Preview",
                show: true
              });
            }
          });
        },
        this.selectedItem.extension !== "jpg"
          ? `image/${this.selectedItem.extension}`
          : "image/jpeg"
      );
    },
    /**
     * 显示相应的模块
     * @param modalName
     */
    showModal(modalName) {
      // this.showCropper = false;
      EventBus.$emit("showCropper", false);
      this.$store.commit("fm/modal/setModalState", {
        modalName,
        show: true
      });
    },
    handleClose(done) {
      done();
      // this.showCropper = false;
      EventBus.$emit("showCropper", false);
      // this.$confirm("确认关闭?", "提示", {
      //   confirmButtonText: "确定",
      //   cancelButtonText: "取消",
      //   type: "warning"
      // })
      //   .then(result => {
      //     if (result == "confirm") {
      //       done();
      //     }
      //   })
      //   .catch(() => {});
    }
  }
};
</script>
<style lang="scss" scoped>
.cropper {
  .title {
    display: inline-block;
    overflow: hidden;
    text-overflow: ellipsis;
    width: -25%;
    width: 90%;
    white-space: nowrap;
    .text-truncate {
      padding-left: 1rem;
      color: #6c757d;
    }
  }

  ::v-deep .el-col {
    .cropper-preview {
      width: 197px;
      height: 200px;
      margin-bottom: 1rem;
      overflow: hidden;
      img {
        min-width: 0px;
        min-height: 0px;
        max-width: none;
        max-height: none;
      }
    }

    .cropper-data {
      padding-left: 1rem;
      padding-right: 1rem;
      height: 245px;
      overflow: auto;
      &::-webkit-scrollbar {
        width: 5px;
      }
      &::-webkit-scrollbar-track {
        background-color: #f5f5f5;
        -webkit-box-shadow: inset 0 0 3px rgba(0, 0, 0, 0.1);
        border-radius: 5px;
      }
      &::-webkit-scrollbar-thumb {
        background-color: rgba(140, 199, 181, 0.8);
        border-radius: 5px;
      }
      &::-webkit-scrollbar-button {
        display: none;
        background-color: #eee;
      }
      &::-webkit-scrollbar-corner {
        background-color: black;
      }
      .el-form {
        .el-form-item {
          margin-bottom: 0.5rem;
          position: relative;
          width: 100%;
          .el-form-item__label {
            width: 28% !important;
            padding: 0.25rem 0.2rem;
            font-size: 0.875rem;
            line-height: 1.5;
            border-radius: 0.2rem;
            border-top-right-radius: 0;
            border-bottom-right-radius: 0;
            text-align: left;
            white-space: nowrap;
            background-color: #e9ecef;
            border: 1px solid #ced4da;
            &::before {
              content: "";
            }
          }
          .el-form-item__content {
            line-height: 0;
            display: flex;
            flex-wrap: wrap;
            align-items: stretch;
            width: 72%;
            .el-input {
              width: 71%;
              .el-input__inner {
                height: 31px;
                border-radius: 0 !important;
              }
            }
            .unit {
              display: inline-block;
              background-color: #e9ecef;
              border: 1px solid #ced4da;
              line-height: 1.2;
              padding: 0.4rem 0.4rem;
              width: 17%;
              border-top-right-radius: 5px;
              border-bottom-right-radius: 5px;
            }
          }
        }
        & > .el-form-item:nth-child(6),
        & > .el-form-item:last-child {
          .el-input {
            width: 99%;
          }
        }
      }
    }
    .check {
      cursor: pointer;
      color: #fff;
      background-color: #17a2b8;
      border: 0px solid #17a2b8;
      padding: 0.25rem 0.5rem;
      font-size: 0.875rem;
      line-height: 1.5;
      border-radius: 0.2rem;
      display: block;
      width: 100%;
      border-style: none;
      margin-bottom: 2px;
      &:hover {
        color: #fff;
        background-color: #138496;
        border: 0px solid #117a8b;
      }
    }
  }
  .dialog-footer {
    .btn-group {
      .el-button {
        cursor: pointer;
        padding: 0.675rem 0.75rem;
        color: #fff;
        background-color: #17a2b8;
        border-color: #17a2b8;
        font-weight: 400;
        text-align: center;
        vertical-align: middle;
        &:hover {
          color: #fff;
          background-color: #138496;
          border-color: #117a8b;
        }
      }

      & > .el-button:not(:last-child) {
        border-top-right-radius: 0;
        border-bottom-right-radius: 0;
        margin-left: 10px;
      }
      & > .el-button:not(:first-child) {
        border-top-left-radius: 0;
        border-bottom-left-radius: 0;
        margin: 0;
      }
    }
    .info,
    .save {
      padding: 0 0.75rem;
      height: 38px;
    }
    .info {
      margin-left: 10px;
      color: #fff;
      background-color: #17a2b8;
      border-color: #17a2b8;
      &:hover {
        color: #fff;
        background-color: #138496;
        border-color: #117a8b;
      }
    }
    .back {
      .el-button {
        color: #212529;
        background-color: #f8f9fa;
        border-color: #f8f9fa;
        padding-left: 1rem;
        padding-right: 1rem;
        &:hover {
          background-color: #e2e6ea;
        }
      }
    }
  }
}
::v-deep .el-dialog .el-dialog__body {
  padding: 0px 1px;
  .cropper-container.cropper-bg {
    width: 100% !important;
  }
}
</style>