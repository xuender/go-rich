import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { Customer } from '../../api/customer';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.page.html',
  styleUrls: ['./customer.page.scss'],
})
export class CustomerPage implements OnInit {
  customer: Customer
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.customer = this.navParams.get('customer')
    if(!this.customer.extend){
      this.customer.extend = {}
    }
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
