import { Component, OnInit } from '@angular/core';
import { CustomerService } from '../../api/customer.service';
import { pull, pick, omit } from 'lodash-es'
import { ModalController } from '@ionic/angular';
import { FileUploader, FileItem } from 'ng2-file-upload';
const KEYS = ['name', 'phone', 'note']

@Component({
  selector: 'app-upload',
  templateUrl: './upload.page.html',
  styleUrls: ['./upload.page.scss'],
})
export class UploadPage implements OnInit {
  pro: any = { name: 1, phone: 2, note: 3 }
  exp: any[] = []
  isEdit: boolean = false
  uploader: FileUploader = new FileUploader({ url: `${CustomerService.URL}/api/c/up` });


  constructor(
    public modalCtrl: ModalController,
    public customer: CustomerService
  ) {
    this.uploader.onCompleteItem = (item: FileItem, r: string, status: number) => {
      if (status !== 200) {
        // TOOD 服务器错误
        return
      }
      this.customer.load()
      this.modalCtrl.dismiss()
    }
  }

  ngOnInit() {
    this.customer.promap().then(a => {
      this.pro = a
      this.exp = []
      for (let k in omit(this.pro, KEYS)) {
        this.exp.push({
          name: k,
          value: this.pro[k],
        })
      }
      this.pro = pick(this.pro, KEYS)
    })
  }
  remove(e: any) {
    pull(this.exp, e)
  }
  cancel() {
    this.modalCtrl.dismiss();
  }
  save() {
    this.isEdit = false
    for (let e of this.exp) {
      this.pro[e.name] = e.value
    }
    this.customer.savePromap(this.pro)
      .subscribe(() => this.modalCtrl.dismiss())
  }
}
