<template>
  <div class="fm-notification">
    <transition-group name="notify">
      <div
        class="fm-notification-item"
        role="alert"
        v-for="(notification, index) in notifications"
        v-bind:class="`fm-${notification.status}`"
        v-bind:key="`notify-${index}`"
      >{{ notification.message }}</div>
    </transition-group>
  </div>
</template>

<script>
import EventBus from "../../eventBus";

export default {
  name: "notification",
  data() {
    return {
      notifications: []
    };
  },
  mounted() {
    EventBus.$on("addNotification", ({ status, message }) =>
      this.addNotification(status, message)
    );
  },
  methods: {
    /**
     * 显示新通知
     * @param status
     * @param message
     */
    addNotification(status, message) {
      this.notifications.push({
        status,
        message
      });
      // 超时后关闭
      setTimeout(() => {
        this.notifications.shift();
      }, 3000);
    }
  }
};
</script>

<style lang="scss" scoped>
.fm-notification {
  position: absolute;
  right: 1rem;
  bottom: 0;
  z-index: 9999;
  width: 350px;
  display: block;
  transition: opacity 0.4s ease;
  overflow: auto;

  .fm-notification-item {
    padding: 0.75rem 1.25rem;
    margin-bottom: 1rem;
    border: 1px solid;
    border-radius: 0.25rem;
    .fm-error {
      color: white;
      background-color: #dc3545;
      border-color: #dc3545;
    }

    .fm-danger {
      color: #dc3545;
      background-color: white;
      border-color: #dc3545;
    }

    .fm-warning {
      color: #ffc107;
      background-color: white;
      border-color: #ffc107;
    }

    .fm-success {
      color: #28a745;
      background-color: white;
      border-color: #28a745;
    }

    .fm-info {
      color: #17a2b8;
      background-color: white;
      border-color: #17a2b8;
    }
  }

  .notify-enter-active {
    transition: all 0.3s ease;
  }
  .notify-leave-active {
    transition: all 0.8s ease;
  }
  .notify-enter,
  .notify-leave-to {
    opacity: 0;
  }
}
</style>
