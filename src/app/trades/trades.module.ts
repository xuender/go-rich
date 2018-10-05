import { NgModule } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';

import { IonicModule } from '@ionic/angular';

import { TradesPage } from './trades.page';
import { TradePageModule } from './trade/trade.module';
import { PipeModule } from '../pipe/pipe.module';

const routes: Routes = [
  {
    path: '',
    component: TradesPage
  }
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    RouterModule.forChild(routes),
    TradePageModule,
    PipeModule,
  ],
  declarations: [TradesPage],
  providers: [DatePipe,]
})
export class TradesPageModule { }
