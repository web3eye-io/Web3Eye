import { uid } from 'quasar'
import { MenuItem } from './types'

const HomePageBreadcrumbs = {
  menuId: uid(),
  icon: 'home',
  label: 'Home',
  target: '/',
  caption: 'Home Page',
  sectionBegin: true,
  level: 0,
  children: []
} as MenuItem

export {
  HomePageBreadcrumbs
}
