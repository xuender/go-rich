import { Component, OnInit, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { URL } from '../../api/init'
import { ModalController, ActionSheetController } from '@ionic/angular';
import { XlsxPage } from './xlsx/xlsx.page';
import { Xlsx } from '../../api/xlsx';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-xlsxes',
  templateUrl: './xlsxes.page.html',
  styleUrls: ['./xlsxes.page.scss'],
})
export class XlsxesPage implements OnInit {
  xlsxes: Observable<Xlsx[]>
  constructor(
    private http: HttpClient,
    private actionSheetCtrl: ActionSheetController,
    private ngZone: NgZone,
    private modalCtrl: ModalController
  ) {
    this.load()
  }
  load() {
    this.xlsxes = this.http.get<Xlsx[]>(`${URL}/api/xlsxes`)
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
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.post<Xlsx>(`${URL}/api/xlsxes`, d.data)
          .subscribe(() => {
            this.load()
          })
      }
    })
    return await modal.present()
  }
  async update(x: Xlsx) {
    const modal = await this.modalCtrl.create({
      component: XlsxPage,
      componentProps: { xlsx: Object.assign({}, x) },
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.put<Xlsx>(`${URL}/api/xlsxes/${x.id}`, d.data)
          .subscribe(nx => {
            Object.assign(x, nx)
          })
      }
    })
    return await modal.present()
  }
  async del(x: Xlsx) {
    const actionSheet = await this.actionSheetCtrl.create({
      header: `确认删除Excel定义 [ ${x.name} ] ?`,
      buttons: [{
        text: '删除',
        role: 'destructive',
        icon: 'trash',
        handler: () => {
          this.http.delete(`${URL}/api/xlsxes/${x.id}`)
            .subscribe((r) => {
              this.load()
              // 强制刷新
              this.ngZone.run(() => false)
            })
        }
      }, {
        text: '取消',
        icon: 'close',
        role: 'cancel',
      }]
    });
    await actionSheet.present();
  }
}
