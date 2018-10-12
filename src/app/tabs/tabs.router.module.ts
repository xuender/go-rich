import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { TabsPage } from './tabs.page';
import { ItemsPage } from '../items/items.page';
import { TradesPage } from '../trades/trades.page';
import { SettingPage } from '../setting/setting.page';
import { CustomersPage } from '../customers/customers.page';
import { AboutPage } from '../setting/about/about.page';

const routes: Routes = [
  {
    path: 'tabs',
    component: TabsPage,
    children: [
      {
        path: '',
        redirectTo: '/tabs/(customers:customers)',
        pathMatch: 'full',
      },
      {
        path: 'customers',
        outlet: 'customers',
        component: CustomersPage
      },
      {
        path: 'trades',
        outlet: 'trades',
        component: TradesPage,
      },
      {
        path: 'items',
        outlet: 'items',
        component: ItemsPage
      },
      {
        path: 'setting',
        outlet: 'setting',
        component: SettingPage,
      },
      {
        path: 'about',
        outlet: 'setting',
        component: AboutPage,
      },
    ]
  },
  {
    path: '',
    redirectTo: '/tabs/(customers:customers)',
    pathMatch: 'full'
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class TabsPageRoutingModule { }
