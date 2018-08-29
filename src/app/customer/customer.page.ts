import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { map } from 'lodash-es'
import { Customer, Ext} from '../api/customer';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.page.html',
  styleUrls: ['./customer.page.scss'],
})
export class CustomerPage implements OnInit {
  customer: Customer
  constructor(
    public modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.customer = this.navParams.get('customer')
  }
  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    this.modalCtrl.dismiss(this.customer);
  }

  ngOnInit() {
  }

}
