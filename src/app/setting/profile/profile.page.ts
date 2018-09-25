import { ModalController } from '@ionic/angular';
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { URL } from '../../api/init'
import { ProfileService } from '../../api/profile.service';
@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {
  user: any = { id: '', nick: '' }
  constructor(
    private modalCtrl: ModalController,
    private profile: ProfileService,
    private http: HttpClient
  ) {
    this.http.get<any>(`${URL}/api/profile`)
      .subscribe(p => this.user = p)
  }

  ngOnInit() {
  }
  logout() {
    this.profile.logout()
    this.profile.login()
  }
  cancel() {
    this.modalCtrl.dismiss();
  }
  pass() {
    // TODO 修改密码
  }
}
