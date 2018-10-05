import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';

import { IonicModule } from '@ionic/angular';

import { TradePage } from './trade.page';
import { PipeModule } from '../../pipe/pipe.module';

const routes: Routes = [
  {
    path: '',
    component: TradePage
  }
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    RouterModule.forChild(routes),
    PipeModule,
  ],
  declarations: [TradePage]
})
export class TradePageModule {}
