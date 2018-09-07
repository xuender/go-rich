import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';

import { IonicModule } from '@ionic/angular';

import { XlsxesPage } from './xlsxes.page';
import { XlsxPageModule } from './xlsx/xlsx.module';

const routes: Routes = [
  {
    path: '',
    component: XlsxesPage
  }
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    XlsxPageModule,
    RouterModule.forChild(routes)
  ],
  declarations: [XlsxesPage]
})
export class XlsxesPageModule {}
