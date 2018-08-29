import { Component, OnInit } from '@angular/core';
import { CustomerService } from '../../api/customer.service';
import { pull, remove, filter, includes, forEach } from 'lodash'
import { ModalController } from '@ionic/angular';
import { FileUploader, FileItem } from 'ng2-file-upload';
import { Ext } from '../../api/customer';
const KEYS = ['Name', 'Phone', 'Note']

@Component({
  selector: 'app-upload',
  templateUrl: './upload.page.html',
  styleUrls: ['./upload.page.scss'],
})
export class UploadPage implements OnInit {
  pro: any = { Name: 1, Phone: 2, Note: 3 }
  ext: Ext[] = []
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
      this.customer.loadGroup()
      this.modalCtrl.dismiss()
    }
  }

  ngOnInit() {
    this.customer.ext().then(es => {
      console.log(es)
      forEach(filter(es, e => includes(KEYS, e.key)), e => this.pro[e.key] = e.value)
      remove(es, (e: Ext) => includes(KEYS, e.key))
      this.ext = es

      // this.pro = a
      // this.exp = []
      // for (let k in omit(this.pro, KEYS)) {
      //   this.exp.push({
      //     name: k,
      //     value: this.pro[k],
      //   })
      // }
      // this.pro = pick(this.pro, KEYS)
    })
  }
  remove(e: any) {
    pull(this.ext, e)
  }
  cancel() {
    this.modalCtrl.dismiss();
  }
  save() {
    this.isEdit = false
    for (const k in this.pro) {
      this.ext.push({ key: k, value: this.pro[k] })
    }
    this.customer.saveExt(this.ext)
      .subscribe(() => this.modalCtrl.dismiss())
  }
}
