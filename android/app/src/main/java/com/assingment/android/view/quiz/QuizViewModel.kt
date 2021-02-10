package com.assingment.android.view.quiz

import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.assingment.android.model.Data
import com.assingment.android.model.QuizQuestion

class QuizViewModel: ViewModel() {

    val quizQuestions = MutableLiveData<List<QuizQuestion>>()
    var page = 0

    init {
        quizQuestions.value = Data.quizQuestions
    }
}
