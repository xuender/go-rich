import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';

import { IonicModule } from '@ionic/angular';

import { ItemsPage } from './items.page';
import { ItemPageModule } from './item/item.module';
import { PipeModule } from '../pipe/pipe.module';

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
  ],
  declarations: [
    ItemsPage,
  ]
})
export class ItemsPageModule {}
