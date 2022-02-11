// store modules
import tree from './tree/store';
import modal from './modal/store';
import settings from './settings/store';
import manager from './manager/store';
import messages from './messages/store';
// main store
import state from './state';
import mutations from './mutations';
import getters from './getters';
import actions from './actions';
export default {
    namespaced: true,
    modules: {
        settings,
        left: manager,
        right: manager,
        tree,
        modal,
        messages,
    },
    state,
    mutations,
    actions,
    getters,
};