import { Injectable } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { XlsxesPage } from '../setting/xlsxes/xlsxes.page';
import { ProfilePage } from '../setting/profile/profile.page';
import { TagsPage } from '../tags/tags.page';
import { UsersPage } from '../users/users.page';

@Injectable({
  providedIn: 'root'
})
export class SettingService {

  constructor(
    private modalCtrl: ModalController
  ) {
  }
  // Excel定义
  async xlsxes() {
    const modal = await this.modalCtrl.create({
      component: XlsxesPage,
    });
    await modal.present()
    return modal.onDidDismiss()
  }
  // 用户
  async users() {
    const modal = await this.modalCtrl.create({
      component: UsersPage,
    });
    return await modal.present()
  }
  // 账户
  async profile() {
    const modal = await this.modalCtrl.create({
      component: ProfilePage,
    });
    return await modal.present()
  }
  // 标签
  async tags() {
    const modal = await this.modalCtrl.create({
      component: TagsPage,
    });
    await modal.present()
    return await modal.onDidDismiss()
  }
}
