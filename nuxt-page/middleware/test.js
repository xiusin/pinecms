export default function ({ store, route, redirect }) {
    store.commit('SET_TOKEN', route.path)
}