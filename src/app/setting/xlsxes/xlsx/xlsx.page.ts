import { Component } from '@angular/core';
import { NavParams, ModalController } from '@ionic/angular';
import { invert } from 'lodash'

import { Xlsx } from '../../../api/xlsx';
import { Ext } from '../../../api/ext';
import { ExtService } from '../../../api/ext.service';
import { Observable } from 'rxjs';

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
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    this.xlsx.map = invert(this.xlsx.map)
    this.modalCtrl.dismiss(this.xlsx);
  }

}
