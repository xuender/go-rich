import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';

import { IonicModule } from '@ionic/angular';

import { SettingPage } from './setting.page';
import { XlsxesPageModule } from './xlsxes/xlsxes.module';
import { UsersPageModule } from './users/users.module';

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
    RouterModule.forChild(routes)
  ],
  declarations: [SettingPage]
})
export class SettingPageModule {}
