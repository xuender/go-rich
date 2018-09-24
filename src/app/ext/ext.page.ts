import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { pull } from 'lodash'

import { Ext } from './ext';
import { ExtService } from './ext.service';

@Component({
  selector: 'app-ext',
  templateUrl: './ext.page.html',
  styleUrls: ['./ext.page.scss'],
})
export class ExtPage {
  exts: Ext[]
  private service: ExtService
  private code: string
  title: string
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.exts = this.navParams.get('exts')
    this.service = this.navParams.get('service')
    this.code = this.navParams.get('code')
    this.title = `${this.code == 'C' ? '客户' : '商品'}扩展`
  }

  async ok() {
    if (this.code == 'C') {
      this.modalCtrl.dismiss(await this.service.saveByCustomer(this.exts))
    } else {
      this.modalCtrl.dismiss(await this.service.saveByItem(this.exts))
    }
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  remove(e: Ext) {
    pull(this.exts, e)
  }
}
