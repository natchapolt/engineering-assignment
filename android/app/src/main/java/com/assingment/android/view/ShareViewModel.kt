package com.assingment.android.view

import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class ShareViewModel: ViewModel() {
    val name = MutableLiveData<String>()
    var score = 0
}
