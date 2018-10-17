import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { pull } from 'lodash'

import { Item, Batch } from '../item';
import { ObjPage } from '../../api/obj.page';
import { ItemService } from '../item.service';
import { ToMoney } from '../../api/money';

@Component({
  selector: 'app-item',
  templateUrl: './item.page.html',
  styleUrls: ['./item.page.scss'],
})
export class ItemPage extends ObjPage<Item> {
  constructor(
    private itemService: ItemService,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
    if (!this.obj.extend) { this.obj.extend = {} }
    if (!this.obj.tags) { this.obj.tags = [] }
    if (!this.obj.cost) { this.obj.cost = 0 }
    if (!this.obj.batchs) { this.obj.batchs = [] }
    for (let i = 0; i < this.obj.batchs.length; i++) {
      if (!this.obj.batchs[i].costMoney) {
        this.obj.batchs[i] = ToMoney(this.obj.batchs[i])
      }
    }
  }

  addBatch() {
    this.obj.batchs.push({
      cost: 0,
      total: 0,
      inventory: 0,
    })
  }

  remove(b: Batch) {
    pull(this.obj.batchs, b)
  }

  get service() { return this.itemService }

  get itemPrice() {
    return this.obj.price / 100
  }

  set itemPrice(price: number) {
    this.obj.price = Math.ceil(price * 100)
  }

  get itemCost() {
    return this.obj.cost / 100
  }

  set itemCost(cost: number) {
    this.obj.cost = Math.ceil(cost * 100)
  }
}
