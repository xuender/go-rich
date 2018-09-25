import { NgModule } from '@angular/core';
import { IonicModule } from '@ionic/angular';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

import { ItemsPage } from './items.page';
import { PipeModule } from '../pipe/pipe.module';
import { ItemPageModule } from './item/item.module';
import { SelectPageModule } from '../tags/select/select.module';

const routes: Routes = [
  {
    path: '',
    component: ItemsPage
  }
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    RouterModule.forChild(routes),
    ItemPageModule,
    PipeModule,
    SelectPageModule,
  ],
  declarations: [
    ItemsPage,
  ]
})
export class ItemsPageModule { }
