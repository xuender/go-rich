import { Component } from '@angular/core';
import { ModalController, ToastController } from '@ionic/angular';

import { ProfileService } from './profile.service';
@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage {
  user: any = { id: '', nick: '' }
  constructor(
    private modalCtrl: ModalController,
    private profile: ProfileService,
    private toastCtrl: ToastController,
  ) {
    this.profile.profile$.subscribe(p => this.user = p)
  }

  logout() {
    this.profile.logout()
    this.modalCtrl.dismiss();
    this.profile.login()
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  async save() {
    await this.profile.save(this.user).toPromise()
    const toast = await this.toastCtrl.create({ message: '账户信息保存成功', duration: 2000, position: 'middle', })
    toast.present()
  }

  pass() {
    this.profile.pass()
  }
}
