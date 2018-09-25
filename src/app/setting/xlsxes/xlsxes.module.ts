import { NgModule } from '@angular/core';
import { IonicModule } from '@ionic/angular';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

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
export class XlsxesPageModule { }
