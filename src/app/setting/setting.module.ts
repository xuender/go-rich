import { NgModule } from '@angular/core';
import { IonicModule } from '@ionic/angular';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

import { SettingPage } from './setting.page';
import { UsersPageModule } from '../users/users.module';
import { XlsxesPageModule } from './xlsxes/xlsxes.module';
import { AboutPageModule } from './about/about.module';

const routes: Routes = [
  {
    path: '',
    component: SettingPage
  }
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    XlsxesPageModule,
    UsersPageModule,
    RouterModule.forChild(routes),
    AboutPageModule,
  ],
  declarations: [SettingPage]
})
export class SettingPageModule { }
