import { Component } from '@angular/core';
import { ModalController, NavParams, AlertController } from '@ionic/angular';
import { FileUploader, FileItem, Headers } from 'ng2-file-upload';
import { Observable } from 'rxjs';

import { Config } from '../../../api/config'
import { Xlsx } from '../../../setting/xlsxes/xlsx';
import { CustomerService } from '../../../customers/customer.service';
import { SettingService } from '../../../setting/setting.service';
import { XlsxService } from '../xlsx.service';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.page.html',
  styleUrls: ['./upload.page.scss'],
})
export class UploadPage {
  header: Headers = {
    name: 'xlsx',
    value: '',
  }
  headers: Headers[] = [this.header]
  uploader: FileUploader
  select: Xlsx
  xlsxes$: Observable<Xlsx[]>
  constructor(
    public modalCtrl: ModalController,
    public customer: CustomerService,
    public setting: SettingService,
    private xlsxService: XlsxService,
    private navParams: NavParams,
    private alertCtrl: AlertController,
  ) {
    this.xlsxes$ = this.xlsxService.all$
    this.uploader = new FileUploader({
      url: `${Config.URL}/api/${this.navParams.get('url')}`,
      headers: this.headers,
      authToken: `Bearer ${localStorage.getItem("token")}`
    })
    this.uploader.onCompleteItem = (item: FileItem, r: string, status: number) => {
      if (status !== 200) {
        this.alertCtrl.create({
          header: `错误 [ ${status} ]`,
          message: JSON.parse(r).message,
          buttons: ['确定']
        }).then(a => a.present())
        return
      }
      this.modalCtrl.dismiss()
    }
  }

  async setup() {
    await this.setting.xlsxes()
    this.xlsxes$ = this.xlsxService.all$
  }

  cancel() {
    this.modalCtrl.dismiss();
  }
}
