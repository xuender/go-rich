import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';

import { UserService } from '../../api/user.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.page.html',
  styleUrls: ['./users.page.scss'],
})
export class UsersPage implements OnInit {

  constructor(
    private modalCtrl: ModalController,
    public user: UserService,
  ) { }

  ngOnInit() {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }
}
