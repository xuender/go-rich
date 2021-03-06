import { NgModule } from '@angular/core';
import { IonicModule } from '@ionic/angular';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

import { TabsPage } from './tabs.page';
import { ItemsPageModule } from '../items/items.module';
import { TabsPageRoutingModule } from './tabs.router.module';
import { SettingPageModule } from '../setting/setting.module';
import { ContactPageModule } from '../customers/customers.module';
import { TradesPageModule } from '../trades/trades.module';
import { AboutPageModule } from '../setting/about/about.module';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    TabsPageRoutingModule,
    SettingPageModule,
    ContactPageModule,
    ItemsPageModule,
    TradesPageModule,
    AboutPageModule,
  ],
  declarations: [TabsPage]
})
export class TabsPageModule { }
