import {MDCSelect} from '@material/select';
import {MDCDrawer} from "@material/drawer";
import {MDCList} from "@material/list";
import {MDCTopAppBar} from "@material/top-app-bar";

const topAppBar = new MDCTopAppBar.attachTo(document.querySelector('#app-bar'));
const list = new MDCList.attachTo(document.querySelector('.mdc-list'));
const drawer = new MDCDrawer.attachTo(document.querySelector('.mdc-drawer'));

// topAppBar.setScrollTarget(document.getElementById('main-content'));
topAppBar.listen('MDCTopAppBar:nav', () => {
  drawer.open = !drawer.open;
});
