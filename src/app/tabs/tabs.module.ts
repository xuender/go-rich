import { IonicModule } from '@ionic/angular';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { TabsPageRoutingModule } from './tabs.router.module';

import { TabsPage } from './tabs.page';
import { ContactPageModule } from '../customers/customers.module';
import { SettingPageModule } from '../setting/setting.module';
import { HomePageModule } from '../home/home.module';
import { ItemsPageModule } from '../items/items.module';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    TabsPageRoutingModule,
    SettingPageModule,
    ContactPageModule,
    HomePageModule,
    ItemsPageModule,
  ],
  declarations: [TabsPage]
})
export class TabsPageModule { }
