import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { map } from 'lodash-es'
import { Customer, Property } from '../api/customer';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.page.html',
  styleUrls: ['./customer.page.scss'],
})
export class CustomerPage implements OnInit {
  customer: Customer
  properties: Property[]
  constructor(
    public modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.customer = this.navParams.get('customer')
    this.properties = map(this.customer.properties, (v, k) => {
      return {
        key: k,
        value: v,
      }
    })
  }
  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    for (const p of this.properties) {
      this.customer.properties[p.key] = p.value
    }
    this.modalCtrl.dismiss(this.customer);
  }

  ngOnInit() {
  }

}
