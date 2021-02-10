package com.assingment.android

import androidx.arch.core.executor.testing.InstantTaskExecutorRule
import com.assingment.android.view.quiz.QuizViewModel
import junit.framework.TestCase.assertEquals
import org.junit.Test
import junit.framework.TestCase.assertNotNull
import org.junit.Rule
import org.junit.rules.TestRule

class QuizViewModelTest {

    @get:Rule
    var rule: TestRule = InstantTaskExecutorRule()

    @Test
    fun testInitQuizViewModel() {
        val quizViewModel = QuizViewModel()
        assertNotNull(quizViewModel.quizQuestions.value)
        assertEquals(10, quizViewModel.quizQuestions.value?.size)
        assertEquals(0, quizViewModel.page)
    }
}
