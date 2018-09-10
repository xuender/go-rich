import { Component, OnInit } from '@angular/core';
import { XlsxService } from '../../api/xlsx.service';
import { ModalController } from '@ionic/angular';
import { XlsxPage } from './xlsx/xlsx.page';
import { Xlsx } from '../../api/xlsx';

@Component({
  selector: 'app-xlsxes',
  templateUrl: './xlsxes.page.html',
  styleUrls: ['./xlsxes.page.scss'],
})
export class XlsxesPage implements OnInit {
  constructor(
    public xlsx: XlsxService,
    private modalCtrl: ModalController,
  ) {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }
  ngOnInit() {
  }

  async create() {
    const modal = await this.modalCtrl.create({
      component: XlsxPage,
      componentProps: {
        xlsx: {
          name: '新Excel定义',
          map: {},
        }
      }
    });
    modal.onDidDismiss().then(d => this.xlsx.post(d.data))
    return await modal.present()
  }
  async update(x: Xlsx) {
    const modal = await this.modalCtrl.create({
      component: XlsxPage,
      componentProps: { xlsx: Object.assign({}, x) },
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.xlsx.put(d.data)
          .then(nx => {
            Object.assign(x, nx)
          })
      }
    })
    return await modal.present()
  }
}
