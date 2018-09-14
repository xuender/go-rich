import { Injectable } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { XlsxesPage } from '../setting/xlsxes/xlsxes.page';
import { UsersPage } from '../setting/users/users.page';
import { ProfilePage } from '../setting/profile/profile.page';

@Injectable({
  providedIn: 'root'
})
export class SettingService {

  constructor(
    private modalCtrl: ModalController
  ) {
  }
  /**
   * Excel定义
   */
  async xlsxes() {
    const modal = await this.modalCtrl.create({
      component: XlsxesPage,
    });
    return await modal.present()
  }
  /**
   * 用户列表
   */
  async users() {
    const modal = await this.modalCtrl.create({
      component: UsersPage,
    });
    return await modal.present()
  }
  /**
   * 账户
   */
  async profile() {
    const modal = await this.modalCtrl.create({
      component: ProfilePage,
    });
    return await modal.present()
  }
}
