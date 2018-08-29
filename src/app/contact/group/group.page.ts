import { Component, OnInit } from '@angular/core';
import { CustomerService } from '../../api/customer.service';
import { ModalController } from '@ionic/angular';

@Component({
  selector: 'app-group',
  templateUrl: './group.page.html',
  styleUrls: ['./group.page.scss'],
})
export class GroupPage implements OnInit {

  group
  constructor(
    private modalCtrl: ModalController,
    public customer: CustomerService
  ) { }

  select(g: string) {
    this.modalCtrl.dismiss(g);
  }
  cancel() {
    this.modalCtrl.dismiss();
  }

  ngOnInit() {
  }

}
