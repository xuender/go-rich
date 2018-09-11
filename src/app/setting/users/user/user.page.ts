import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { User } from '../../../api/user';

@Component({
  selector: 'app-user',
  templateUrl: './user.page.html',
  styleUrls: ['./user.page.scss'],
})
export class UserPage implements OnInit {
  user: User
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.user = this.navParams.get('user')
  }

  ngOnInit() {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    this.modalCtrl.dismiss(this.user);
  }
}
