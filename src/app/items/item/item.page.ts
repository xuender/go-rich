import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { Item } from '../../api/item';

@Component({
  selector: 'app-item',
  templateUrl: './item.page.html',
  styleUrls: ['./item.page.scss'],
})
export class ItemPage implements OnInit {

  item: Item
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.item = this.navParams.get('item')
  }

  ngOnInit() {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    this.modalCtrl.dismiss(this.item);
  }
}
