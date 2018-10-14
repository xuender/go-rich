import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Item } from '../item';
import { ObjPage } from '../../api/obj.page';
import { ItemService } from '../item.service';

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
