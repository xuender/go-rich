import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Item } from '../item';
import { ItemService } from '../item.service';
import { ObjPage } from '../../api/obj.page';

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
  }

  get service() { return this.itemService }

  get itemPrice() {
    return this.obj.price / 100
  }

  set itemPrice(price: number) {
    this.obj.price = Math.ceil(price * 100)
  }
}
