import Vue from "vue";
import Component from "./Crud.vue";

const Constructor = Vue.extend(Component);

const ComponentExample = options => {
  if (Vue.prototype.$isServer) return;
  const { ...rest } = options;
  const instance = new Constructor({
    propsData: {
      ...rest
    }
  });

  const id = `crudModal`;
  instance.id = id;
  instance.vm = instance.$mount();
  document.body.appendChild(instance.vm.$el);
  instance.vm.visible = true;

  //绑定 success 方法
  instance.vm.$on("success", values => {
    if (instance.success) {
      instance.success(values);
    }
  });
  instance.vm.$on("cancel", () => {
    if (instance.cancel) {
      instance.cancel();
      document.body.removeChild(instance.vm.$el);
    }
  });
  instance.vm.$on("destroy", () => {
    document.body.removeChild(instance.vm.$el);
  });
  return instance.vm;
};

export default ComponentExample;
