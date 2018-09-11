import { Component } from '@angular/core';
import { NavParams, ModalController } from '@ionic/angular';
import { invert } from 'lodash'

import { Xlsx } from '../../../api/xlsx';
import { ExtService } from '../../../api/ext.service';

@Component({
  selector: 'app-xlsx',
  templateUrl: './xlsx.page.html',
  styleUrls: ['./xlsx.page.scss'],
})
export class XlsxPage {
  xlsx: Xlsx
  abc = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('')
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams,
    public ext: ExtService
  ) {
    this.xlsx = this.navParams.get('xlsx')
    this.xlsx.map = invert(this.xlsx.map)
    for (const k in this.xlsx.map) {
      if (this.xlsx.map[k]) {
        this.xlsx.map[k] = parseInt(this.xlsx.map[k]) + 1
      }
    }
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    for (const k in this.xlsx.map) {
      if (!this.xlsx.map[k]) {
        delete this.xlsx.map[k]
      } else {
        this.xlsx.map[k] = this.xlsx.map[k] - 1
      }
    }
    this.xlsx.map = invert(this.xlsx.map)
    this.modalCtrl.dismiss(this.xlsx);
  }

}
