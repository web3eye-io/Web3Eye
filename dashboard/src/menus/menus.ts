import { uid } from 'quasar'
import { MenuItem } from 'src/localstore'

const MainDrawerMenus = [
   {
    menuId: uid(),
    label: 'Overview',
    icon: 'pending',
    target: '/manager',
    level: 0,
    sectionBegin: false,
    children: [
      {
        menuId: uid(),
        label: 'Endpoint MGR',
        caption: '',
        icon: 'perm_identity',
        target: '/manager/overview',
        level: 1,
        sectionBegin: false,
        children: []
      } as MenuItem
    ]
  } as MenuItem,
   {
    menuId: uid(),
    label: 'Task',
    icon: 'pending',
    target: '/manager',
    level: 0,
    sectionBegin: false,
    children: [
      {
        menuId: uid(),
        label: 'Task MGR',
        caption: '',
        icon: 'perm_identity',
        target: '/manager/task',
        level: 1,
        sectionBegin: false,
        children: []
      } as MenuItem
    ]
  } as MenuItem,
   {
    menuId: uid(),
    label: 'Endpoint',
    icon: 'pending',
    target: '/manager',
    level: 0,
    sectionBegin: false,
    children: [
      {
        menuId: uid(),
        label: 'Endpoint MRG',
        caption: '',
        icon: 'perm_identity',
        target: '/manager/endpoint',
        level: 1,
        sectionBegin: false,
        children: []
      } as MenuItem
    ]
  } as MenuItem,
] as Array<MenuItem>

export {
  MainDrawerMenus
}
