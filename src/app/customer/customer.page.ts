import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { Customer } from '../api/customer';
import { ViewController } from '@ionic/core';

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
