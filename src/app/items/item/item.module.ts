import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';
import { IonicModule } from '@ionic/angular';

import { ItemPage } from './item.page';
import { ComponentsModule } from '../../component/components.module';

const routes: Routes = [
  {
    path: '',
    component: ItemPage
  }
]

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    RouterModule.forChild(routes),
    ComponentsModule,
  ],
  declarations: [
    ItemPage,
  ]
})
export class ItemPageModule { }
